import typer
import deploy

def deploy_local(
    upgrade: bool = False,
    chart_dir: str = "../helm-charts/go-server",
) -> None:
    deploy.deploy(
        environment="local",
        debug=True,
        upgrade=upgrade,
        chart_dir=chart_dir,
    )

if __name__ == "__main__":
    typer.run(deploy_local)