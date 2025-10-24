# Substring remover

A simple Go utility to recursively remove a given substring from the names of all files and directories in a specified folder.

Make sure you have installed [Go](https://go.dev/) (version 1.24.4 or higher)

### Usage

Clone the repository and navigate to the project directory:
```
git clone https://bulatzaripow/substring-remover
cd substring-remover
```

Run the tool directly:
```
go run . -dir="/path/to/target/folder" -substr="substring_to_remove"
```

Command line flags:
- `-dir` - path to the root directory (_Required_)
- `-substr` - substring to remove (_Required_)

### Safety

- Always `back up` your data before bulk operations
- The script processes items from deepest to shallowest to avoid path conflicts.
- It skips renaming if the new name would be empty.
- It skips renaming if a file/directory with the new name already exists.