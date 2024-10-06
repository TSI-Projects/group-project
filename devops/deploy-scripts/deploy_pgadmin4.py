import typer
import subprocess

def destroy_pgadmin4(
    environment: str = "local",
) -> None:
    p1 = subprocess.Popen([
        "kubectl", "cnpg", "pgadmin4", f"{environment}-postgres-cluster",
        "--dry-run", "--namespace", environment,
    ], stdout=subprocess.PIPE)

    p2 = subprocess.Popen([
        'kubectl', 'delete', '-f', '-'
    ], stdin=p1.stdout, stdout=subprocess.PIPE)

    p1.stdout.close()
    output = p2.communicate()[0]
    typer.echo(output)

def deploy_pgadmin4(
    environment: str = "local",
    destroy: bool = False,
    port: int = 8888,
) -> None:
    if destroy:
        destroy_pgadmin4(environment)
        return

    try:
        subprocess.call([
            "kubectl", "cnpg", "pgadmin4",
            f"{environment}-postgres-cluster",
            "--namespace", environment
        ])

        subprocess.call([
            "kubectl", "rollout", "status", "deployment",
            f"{environment}-postgres-cluster-pgadmin4",
            "--namespace", environment
        ])

        subprocess.call([
            "kubectl", "port-forward",
            f"deployment/{environment}-postgres-cluster-pgadmin4", f"{port}:80",
            "--namespace", environment
        ])
    except KeyboardInterrupt:
        destroy_pgadmin4(environment)

if __name__ == "__main__":
    typer.run(deploy_pgadmin4)
