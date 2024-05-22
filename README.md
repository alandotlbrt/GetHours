#  GetHours / alebrume | alandotlbrt

## REQUIRED 

### GO & IDE

- before running this code you need an IDE && golang language 

- Linux Terminal && Mac Terminal 
```
sudo apt-get update
sudo apt-get install golang

sudo apt-get update
sudo apt-get install code
```

## HOW TO USE 

### .EXE 

- if the exe doesn't exist you have just to run the main programme : 

```
go run buildexe.go
```

### LAUNCH THE PROGRAM 

- you have just to run the exe 

````
./hours01
````

- **BUT** don't forget to **create** the .**exe** before. 


### LAUNCH EVERY WHERE YOU WANT 


- I'm sorry, for now I haven't found a better method for that, so you have to go into a Go file to change a path, **EXCEPT** if you want to use the program in the same folder where it is located, so follow this method.

#### ON MAC && LINUX

- 1 - go to the directory containing the executable file.

- 2 - Launch the **main.go** using the command:

````
go run . 
`````

- 3 - Choose the second option and copy the path.

- 4 Now, go to GoFolder/MAIN/getHours.go

- 5 ocate the variable named YourPath and paste the copied path.

- 6 Repeat step 2, then choose the first option and enter the desired executable name.

- 7 Do these commands : 
```

sudo cp nameExeFile /usr/local/bin/

export PATH="/usr/local/bin:$PATH"

````

- Restart your IDE.

- Now, you can run the code from anywhere by typing the executable name

#### ON WINDOWS 

- 1 - go to the directory containing the executable file.

- 2 - Launch the **main.go** using the command:

````
go run . 
`````

- 3 - Choose the second option and copy the path.

- 4 Now, go to GoFolder/MAIN/getHours.go

- 5 Locate the variable named YourPath and paste the copied path.

- 6 Repeat step 2, then choose the first option and enter the desired executable name.

- 7 Copy the exact path of the exe in your folder, and do these commands : 
```

copy PathOfYourFile C:\Windows\System32\


setx PATH "%PATH%;C:\Windows\System32\"

````

#### USE IT

- Restart your IDE.

- Now, you can run the code from anywhere by typing the executable name


- **BE CAREFUL : DON'T CHANGE THE LOCATION WHERE YOUR ORIGNAL FOLDER IS OR GO TO STEP 1**

## FEATURES 

- Soon I will try to return every public holiday, I will keep you informed. ( i added but i don't really know if it's work for now )

## MORE INFORMATIONS 

- the language you choose will remain for future executions, by default the language will be english.
- the code is supposed to adapt to the end of the month.
- it is possible that there are some bugs, if this is the case let me know, my contacts will be just below

#### SOURCES 

- [Movable public holidays in France](https://fr.wikipedia.org/wiki/Calcul_de_la_date_de_Pâques_selon_la_méthode_de_Meeus)

#### Authors
Made by [Alan L.](https://zone01normandie.org/git/alebrume)

#### Contacts 


- Discord - alan.lbrt
- Mail - alanlebrument@icloud.com
- Others Projects : [https://zone01normandie.org/git/alebrume](https://zone01normandie.org/git/alebrume)
