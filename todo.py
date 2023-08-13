import time,sys, pyautogui as pag
#this is the logic
global nput

def sprint(z): #the slow print
    for y in z:
        sys.stdout.write(y)
        sys.stdout.flush()
        time.sleep(.05)

#make a list that will hold 1 - 64
def TodoList():#idk kno if this is vaild yet
    newItem = input()
    nput = input()
    tDlist = list(range(str(int(x,64))))#this gone give me hell
    for x in range(tDlist(newItem)):
        x + 1
    
sprint('return 0\n')
tr = True
"Exit" == False
nput = input()

if nput.lower() == "yes":
    while tr:
        print("Todo Listing")
        nput = input()
        print(f'{nput}')
        
    

