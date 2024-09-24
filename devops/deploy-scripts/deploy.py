import typer
import os
import subprocess

def deploy(
    environment: str,
    debug: bool = False,
    upgrade: bool = False,
    chart_dir: str = "../helm-charts/go-server",
) -> None:
    helm_action = "upgrade" if upgrade else "install"

    if (os.path.isabs(chart_dir) == False):
        chart_dir = os.path.abspath(chart_dir)

    subprocess.call([
        "helm", helm_action, environment, chart_dir,
        "--namespace", environment, "--create-namespace",
        "--set-string", f"environment={environment}",
        "--set", f"debug={debug}",
    ])

if __name__ == "__main__":
    typer.run(deploy)