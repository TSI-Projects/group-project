import os

def build_path(relative_path: str | None, default : str | None = None) -> str:
    if relative_path is None:
        if default is None:
            raise ValueError("You must provide either a relative path or a default path.")
        
        relative_path = default
    
    if os.path.isabs(relative_path):
        return relative_path

    return os.path.abspath(relative_path)