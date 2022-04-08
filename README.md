# FilesDir
***

## Introduction
***

FilesDIR is a module that allows you to search for files in a folder and its subfolders.
It allows you to search for a word, equal, starting or ending with a keyword in all the files found.
You can also filter files by extension and with or without the capital letters.

## Usage
***

### Installation 
For use FilesDIR, you have to add the path of the executable `FilesDIR.exe` by creating a new environment variable `FilesDIR`   

### Use

Open the desired folder, then, in the explorer, launch the FilesDIR program by giving it its arguments.

required arguments:   
`FilesDIR` : This is the main executable of the module.   

possible arguments :   
`-mode=%` : It allows you to choose the keyword you want to search for.   
Ex: `-mode==` Search exact word in file name.    
Ex: `-mode=%` Search word is contains in file name.    
Ex: `-mode=^` the name of the file starts with the searched word.    
Ex: `-mode=$` the name of the file ends with the searched word.    

`-word=%` : Is the word key for search.    
Ex: `-word=test` 

`-ext=%` : Is the extension of file.    
Ex: `-ext=xlsx` 

`-poolsize=%` : Is the number of max threads.     
Ex: `-poolsize=100`   

`-maj` : Allows if set to take into account uppercase letters.     
`-xl` : Allows data to be exported to an Excel file.   

`-devil` : This is the extreme speed mode but it is not 100% reliable and will slow down the pc.   
`-s` : This id the 'Silent mode', he hide all the messages and various actions.   
`-b` : For active the blacklist folder, add folder name in file generated on : %USER%\FilesDIR\blacklist\...     
you can add a text file to the name of the searched keyword or simply add folder names in the `__ALL__`.txt file   
