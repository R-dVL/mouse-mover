import pyautogui
import time

def moverRaton():
    pyautogui.dragRel(0, 100, duration = 1)
    time.sleep(30)
    pyautogui.dragRel(0,-100, duration = 1)
    time.sleep(30)

if __name__ == "__main__":
    print("Taluego...")
    while True:
        moverRaton()