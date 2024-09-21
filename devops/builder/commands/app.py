from __future__ import annotations
from cleo.commands.command import Command
from cleo.helpers import option

class AppInstallCommand(Command):
    name = "app:install"
    description = "Global command to deploy the whole application to the Kubernetes cluster."
    options = [
        option(
            long_name="upgrade",
            short_name="u",
            description="Upgrade the Helm chart.",
            flag=True,
        )
    ]

    def handle(self):
        if self.call("dependencies:install"):
            return 1

        arguments: list = []

        if self.option("upgrade"):
            arguments.append("--upgrade")

        argumentsString = ",".join(arguments)

        if self.call("postgres-cluster:install", argumentsString):
            return 1
        
        if self.call("go-server:install", argumentsString):
            return 1
        
        return 0
    
class AppUninstallCommand(Command):
    name = "app:uninstall"
    description = "Global command to uninstall the whole application from the Kubernetes cluster."

    def handle(self):
        if self.call("postgres-cluster:uninstall"):
            return 1
        
        if self.call("go-server:uninstall"):
            return 1
        
        return 0