from __future__ import annotations
from cleo.commands.command import Command
from cleo.helpers import argument, option
import subprocess
import sys

class DependenciesInstallCommand(Command):
    name = "dependencies:install"
    description = "Installs the dependencies required for the application stack."
    
    def install_cloudnative_pg(self) -> None:
        subprocess.run([
            "helm", "install", "cnpg", "cloudnative-pg",
            "--repo", "https://cloudnative-pg.github.io/charts",
            "--namespace", "cnpg-system", "--create-namespace",
        ])

    def install_nginx_ingress_controller(self) -> None:
        subprocess.run([
            "helm", "install", "ingress-nginx", "ingress-nginx",
            "--repo", "https://kubernetes.github.io/ingress-nginx",
            "--namespace", "ingress-nginx", "--create-namespace"
        ])

    def install_reflector(self) -> None:
        try:
            subprocess.run([
                "helm", "install", "emberstack", "reflector",
                "--repo", "https://emberstack.github.io/helm-charts",
                "--namespace", "reflector", "--create-namespace"
            ], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL, check=True)

            self.info("Reflector installed or upgraded successfully.")
        except subprocess.CalledProcessError as e:
            self.line_error(f"An error occurred: {e}")

    def install_jspolicy(self) -> None:
        try:
            subprocess.run([
                "helm", "install", "jspolicy", "jspolicy",
                "--repo", "https://charts.loft.sh",
                "--namespace", "jspolicy", "--create-namespace"
            ], stdout=subprocess.DEVNULL, check=True)

            self.info("JSPolicy installed or upgraded successfully.")
        except subprocess.CalledProcessError as e:
            self.line_error(f"An error occurred: {e}")

    def handle(self):
        self.install_cloudnative_pg()
        self.install_nginx_ingress_controller()
        # self.install_reflector()
        # self.install_jspolicy()

__all__ = [
    "InstallDependenciesCommand"
]