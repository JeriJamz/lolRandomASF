import time,sys, pyautogui as pag
#this is the logic
#this is now an app
global nput

def sprint(z): #the slow print
    for y in z:
        sys.stdout.write(y)
        sys.stdout.flush()
        time.sleep(.05)

def Kfprint(a):
    for b in a:
        sys.stdout.write(b)
        sys.stdout.flush()
        time.sleep(.099)


#make a list that will hold 1 - 64

class ToDo:
    def __init__(self, TodoList):
        self.TodoList
        self.click
        self.clickNdrag
        self.UiFunc
        self.Commands

    def TodoList():#idk kno if this is vaild yet
        newItem = input()#working on a print and idk if this takes the strings or just the numbers yet
        nput = input()#dont think i need this
        tDlist = list(range(str(int(x,64))))#this gone give me hell
        for x in range(tDlist(newItem)):#this is the algo to store the item for the day. Max is 64 for the day
            x + 1

    def Click():#from here down are "features ima add"
        pass

    def clickNdrag():
        pass

    def UiFunc():
        pass

    def Commands():
        pass
    
Kfprint("Do you want to access the to do list?\n")

tr = True

nput = input()

if nput.lower() == "yes":
    while tr:
        print("Todo Listing")
        nput = input()
        print(f'{nput}')
        if nput == '<Exit>':#well now I know how to add input for the features
            sprint("End of loop\n")
            tr = False
        elif nput.lower() == 'Show Todo':
            print(nput)
else:
    sprint("have a good \n")
    

