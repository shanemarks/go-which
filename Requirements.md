Our general plan is to implement the "which" command in GO. It finds where a executable is your systems PATH is located, and outputs its full path. 

The general solution involves:

1) Breaking up all directories on your PATH environment variable and creating a list
2) Scanning each one for a file that matches your query, and storing it
3) printing out all matches

PATH Search: It scans the list of directory paths stored in your system's PATH environment variable andFile Matching: It finds and returns the full file path of the executable program that runs when you type that command.