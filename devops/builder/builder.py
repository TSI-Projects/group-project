from __future__ import annotations
from cleo.application import Application
from commands.postgres_cluster import PostgresClusterInstallCommand, PostgressClusterUninstallCommand
from commands.dependencies import DependenciesInstallCommand
from commands.go_server import GoServerInstallCommand, GoServerUninstallCommand
from commands.app import AppInstallCommand, AppUninstallCommand

app = Application()
app.add(DependenciesInstallCommand())
app.add(PostgressClusterUninstallCommand())
app.add(PostgresClusterInstallCommand())
app.add(GoServerInstallCommand())
app.add(GoServerUninstallCommand())
app.add(AppInstallCommand())
app.add(AppUninstallCommand())

if __name__ == "__main__":
    app.run()