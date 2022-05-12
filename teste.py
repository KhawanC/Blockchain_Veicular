import pyautogui, socket, time

sk = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
sk.connect(("8.8.8.8", 80))

pyautogui.hotkey('ctrl', 'shift', 't')
time.sleep(0.5)
pyautogui.write(sk.getsockname()[0])