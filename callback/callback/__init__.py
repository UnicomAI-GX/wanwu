import logging

from flasgger import Swagger
from flask import Flask

from configs.config import load_config
from extensions.minio import init_minio
from extensions.redis import init_redis
from utils.response import register_error_handlers
from utils.router_trace import add_request_tracing


def create_app():
    app = Flask(__name__)
    app.logger.setLevel(logging.INFO)
    app.config["SWAGGER"] = {"openapi": "3.0.1"}

    # init config
    load_config()

    # init redis
    init_redis()

    # init minio
    init_minio()

    # 初始化 swagger
    Swagger(app)

    # 添加路由追踪
    add_request_tracing(app)

    # 注册异常处理
    register_error_handlers(app)

    # 注册蓝图
    from callback.routes.hello import hello_demo

    app.register_blueprint(hello_demo, url_prefix="/v1")

    return app
