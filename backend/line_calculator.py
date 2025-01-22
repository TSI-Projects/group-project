import os

# Set of file extensions to include (customize as needed)
extensions_to_count = {'.go'} 

def count_lines_in_file(file_path):
    """Count non-empty lines in a single file."""
    try:
        with open(file_path, 'r', encoding='utf-8', errors='ignore') as f:
            return sum(1 for line in f if line.strip())
    except Exception as e:
        print(f"Error reading {file_path}: {e}")
        return 0

def count_lines_in_directory(root_dir):
    total_lines = 0
    # Iterate through directory tree
    for dirpath, dirnames, filenames in os.walk(root_dir):
        # Optionally skip certain directories:
        # if 'venv' in dirnames:
        #     dirnames.remove('venv')  # Skip virtual environments
        for filename in filenames:
            ext = os.path.splitext(filename)[1]
            if ext.lower() in extensions_to_count:
                file_path = os.path.join(dirpath, filename)
                lines = count_lines_in_file(file_path)
                total_lines += lines
                print(f"{file_path}: {lines} lines")
    return total_lines

if __name__ == "__main__":
    # Set the directory you want to analyze; default is the current directory
    directory_to_scan = '.'
    total = count_lines_in_directory(directory_to_scan)
    print(f"\nTotal lines of code: {total}")