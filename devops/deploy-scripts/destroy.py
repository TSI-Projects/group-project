import typer
import os
import subprocess

def destroy(
    environment: str,
) -> None:
    subprocess.call([
        "helm", "uninstall", environment,
        "--namespace", environment
    ])

    subprocess.call([
        'kubectl', 'delete', 'namespace', environment
    ])

if __name__ == "__main__":
    typer.run(destroy)