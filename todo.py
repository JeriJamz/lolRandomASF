import time,sys, pyautogui as pag
#this is the logic
global nput

nput = input()

def sprint(z): #the slow print
    for y in z:
        sys.stdout.write(y)
        sys.stdout.flush()
        time.sleep(.05)

#make a list that will hold 1 - 64
def TodoList():#idk kno if this is vaild yet
    newItem = input()
    tDlist = list(range(str(int(x,64))))
    for x in range(tDlist(newItem)):
        x = newItem

sprint('return 0')

tr = True
"<Exit>" == False

while nput != "<Exit>":
    TodoList()
    

