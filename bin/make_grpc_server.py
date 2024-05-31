import os
import sys
import click
from jinja2 import Environment, FileSystemLoader, select_autoescape

BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
PROJECT = 'spaceone'
PROTO_DIR = os.path.join(BASE_DIR, 'proto')
TEMPLATE_PATH = os.path.join(os.path.dirname(__file__), '../template/gateway')
JINJA_ENV = Environment(
    loader=FileSystemLoader(searchpath=TEMPLATE_PATH),
    autoescape=select_autoescape()
)

def _error(msg):
    print()
    print('[ERROR] %s' % (msg))
    print()
    sys.exit(1)

def _get_services_from_target(target):
    _services = []

    api_path_in_proto = os.path.join(PROTO_DIR, PROJECT, 'api')

    # Get all services name from proto path
    for service in os.listdir(api_path_in_proto):
        service_path_in_proto = os.path.join(api_path_in_proto, service)
        if os.path.isdir(service_path_in_proto):
            _services.append(service)

    if target == 'all':
        return _services
    elif target in _services:
        return [target]
    else:
        _error(f"Target({target}) is not found.")

def get_handlers(service):
    handlers = []
    _services = _get_services_from_target(service)
    for file in os.listdir(f'{BASE_DIR}/dist/gateway/{PROJECT}/api/{service}/v1'):
        if file.endswith(".pb.gw.go"):
            with open(f'{BASE_DIR}/dist/gateway/{PROJECT}/api/{service}/v1/{file}') as _f:
                context = _f.readlines()

                for ctx in context:
                    if ctx.find("HandlerFromEndpoint(") > 0:
                        handlers.append(ctx.split(' ')[1].split('(')[0])
    return handlers


@click.command()
@click.argument('target', default='all')
def make_grpc_gateway_server(**params):
    """
    Cloudforet gRPC Gateway Builder\n
    TARGET, 'all or specific service. (core, identity, inventory, etc.)'
    """

    services = _get_services_from_target(params['target'])
    handlers = {}
    for service in services:
        handlers[service] = get_handlers(service)

    template = JINJA_ENV.get_template('gateway_server.tmpl')
    content = template.render(
        handlers=handlers,
        services=sorted(services),
        project=PROJECT,
        version='v1'
    )

    with open(f'{BASE_DIR}/dist/gateway/{PROJECT}/api/gateway_server.go', 'w') as f:
        f.write(content)


if __name__ == '__main__':
    make_grpc_gateway_server()
