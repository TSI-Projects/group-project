import typer
import destroy

def destroy_local() -> None:
    destroy.destroy(environment="local")

if __name__ == "__main__":
    typer.run(destroy_local)