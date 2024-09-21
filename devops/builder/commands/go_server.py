from __future__ import annotations
from cleo.commands.command import Command
from cleo.helpers import option
from support.path import build_path

import support.constants as constants
import shutil
import subprocess

HELM_CHART_DIR = "../helm-charts/go"
RELEASE_NAME = "production-go-server"

class GoServerInstallCommand(Command):
    name = "go-server:install"
    description = "Installs a Go server to the Kubernetes cluster."
    options = [
        option(
            long_name="upgrade",
            short_name="u",
            description="Upgrade the Go server.",
            flag=True,
        )
    ]

    def handle(self):
        if shutil.which("helm") is None:
            self.line_error("helm could not be found, please install it.")
            return 1
        
        helm_chart_dir = build_path(constants.GO_SERVER_HELM_CHART_DIR)
        
        is_upgrade = bool(self.option("upgrade"))
        helm_action = "upgrade" if is_upgrade else "install"

        subprocess.run([
            "helm", helm_action, constants.GO_SERVER_RELEASE_NAME, helm_chart_dir,
            "--namespace", constants.NAMESPACE, "--create-namespace",
            "--set-string", f"fullnameOverride={constants.GO_SERVER_FULLNAME_OVERRIDE}",
            "--set-string", f"databaseEnv.postgresClusterName={constants.POSTGRES_CLUSTER_FULLNAME_OVERRIDE}",
        ])
        
        return 0
    
class GoServerUninstallCommand(Command):
    name = "go-server:uninstall"
    description = "Destroys a Go server in the Kubernetes cluster."

    def handle(self):
        if shutil.which("helm") is None:
            self.line_error("helm could not be found, please install it.")
            return 1

        subprocess.run([
            "helm", "uninstall", constants.GO_SERVER_RELEASE_NAME,
            "--namespace", constants.NAMESPACE
        ])

        return 0

__all__ = [
    "GoServerInstallCommand",
    "GoServerUninstallCommand"
]