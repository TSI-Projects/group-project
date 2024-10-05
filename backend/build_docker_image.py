import typer
import subprocess

def build_docker_image(
    name: str = "simple_go_server",
    tag: str = "latest",
    port: int = 8000,
    dockerfile_dir: str = ".",
    source_dir: str = ".",
    maingo_file_path: str = "./cmd/app/main.go",
) -> None:
    subprocess.call([
        "docker", "build", dockerfile_dir,
        "-t", f"{name}:{tag}",
        "--build-arg", f"GO_SERVER_PORT={port}",
        "--build-arg", f"GO_SERVER_DIR={source_dir}",
        "--build-arg", f"GO_SERVER_MAIN_PATH={maingo_file_path}",
    ])

if __name__ == "__main__":
    typer.run(build_docker_image)