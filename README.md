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
For use FilesDIR, you can simply launch it from anywhere, but it is better to add it to the environment variables.   

To do this, add the path of the folder containing the module to the environment variables `PATH`:   
![var1](https://github.com/LaM0uette/FilesDIR/blob/V0.2/docs/img/var1.PNG)
![var2](https://github.com/LaM0uette/FilesDIR/blob/V0.2/docs/img/var2.PNG)
![var3](https://github.com/LaM0uette/FilesDIR/blob/V0.2/docs/img/var3.PNG)
![var4](https://github.com/LaM0uette/FilesDIR/blob/V0.2/docs/img/var4.PNG)
![var5](https://github.com/LaM0uette/FilesDIR/blob/V0.2/docs/img/var5.PNG)

Then, you have to add the path of the executable `FilesDIR.exe` by creating a new environment variable `FilesDIR`:   
![var6](https://github.com/LaM0uette/FilesDIR/blob/V0.2/docs/img/var6.PNG)

### Use

Open the desired folder, then, in the explorer, launch the FilesDIR program by giving it its arguments.
![fd1](https://github.com/LaM0uette/FilesDIR/blob/V0.2/docs/img/fd1.PNG)

required arguments:   
`FilesDIR` : This is the main executable of the module.   
`-r` : Is required for run app in CLI mode.    

possible arguments :   
`-mode=%` : It allows you to choose the keyword you want to search for.   
Ex: `-mode==` Search exact word in file name.    
Ex: `-mode=%` Search word is contains in file name.    
Ex: `-mode=^` the name of the file starts with the searched word.    
Ex: `-mode=$` the name of the file ends with the searched word.    

`-word=%` : Is the word key for search.    
Ex: `-word=test` Search `test` in file name.    

`-ext=%` : Is the extension of file.    
Ex: `-ext=xlsx` Search all files with `.xlsx` extension    

`-maj` : Allows if set to take into account uppercase letters.    
