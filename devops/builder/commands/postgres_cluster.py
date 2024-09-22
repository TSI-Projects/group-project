from __future__ import annotations
from cleo.commands.command import Command
from cleo.helpers import option
from support.path import build_path

import support.constants as constants
import shutil
import subprocess

HELM_CHART_DIR = "../helm-charts/postgres"
RELEASE_NAME = "production-postgres-cluster"

class PostgresClusterInstallCommand(Command):
    name = "postgres-cluster:install"
    description = "Installs a PostgreSQL cluster to the Kubernetes cluster."
    options = [
        option(
            long_name="upgrade",
            short_name="u",
            description="Upgrade the PostgreSQL cluster.",
            flag=True,
        )
    ]

    def handle(self):
        if shutil.which("helm") is None:
            self.line_error("helm could not be found, please install it.")
            return 1
        
        helm_chart_dir = build_path(constants.POSTGRES_HELM_CHART_DIR)
        
        is_upgrade = bool(self.option("upgrade"))
        helm_action = "upgrade" if is_upgrade else "install"

        subprocess.run([
            "helm", helm_action, constants.POSTGRES_CLUSTER_RELEASE_NAME, helm_chart_dir,
            "--namespace", constants.NAMESPACE, "--create-namespace",
            "--set-string", f"fullnameOverride={constants.POSTGRES_CLUSTER_FULLNAME_OVERRIDE}"
        ])
        
        return 0

class PostgressClusterUninstallCommand(Command):
    name = "postgres-cluster:uninstall"
    description = "Destroys a PostgreSQL cluster in the Kubernetes cluster."

    def handle(self):
        if shutil.which("helm") is None:
            self.line_error("helm could not be found, please install it.")
            return 1

        subprocess.run([
            "helm", "uninstall", constants.POSTGRES_CLUSTER_RELEASE_NAME,
            "--namespace", constants.NAMESPACE
        ])

        return 0

__all__ = [
    "PostgresClusterInstallCommand",
    "PostgressClusterDestroyCommand"
]