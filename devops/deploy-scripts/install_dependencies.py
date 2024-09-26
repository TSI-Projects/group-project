import typer
import subprocess

def install_dependencies(upgrade: bool = False) -> None:
    helm_action = "upgrade" if upgrade else "install"

    subprocess.call([
        "helm", helm_action, "ingress-nginx", "ingress-nginx",
        "--repo", "https://kubernetes.github.io/ingress-nginx",
        "--namespace", "ingress-nginx", "--create-namespace"
    ])

    subprocess.call([
        "helm", helm_action, "cnpg", "cloudnative-pg",
        "--repo", "https://cloudnative-pg.github.io/charts",
        "--namespace", "cnpg-system", "--create-namespace",
    ])

if __name__ == "__main__":
    typer.run(install_dependencies)