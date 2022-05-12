import pyautogui, socket, time

#sk = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
#sk.connect(("8.8.8.8", 80))
#pyautogui.write(sk.getsockname()[0])

#Estabelecer conexão com a rede e subir o chaincode para a rede
pyautogui.hotkey('ctrl', 'shift', 't')
time.sleep(1)
pyautogui.write('cd Área\ de\ Trabalho\nmiblocknet')
pyautogui.press('enter')
time.sleep(1)
pyautogui.write('docker-compose -f peer-orderer.yamml -f peer-ptb.yaml up -d')
pyautogui.press('enter')
time.sleep(10)
pyautogui.write('./configurechannel.sh ptb.de -c')
pyautogui.press('enter')
time.sleep(15)
pyautogui.write('./configurechaincode.sh install cli0 fabpki 1.0')
pyautogui.press('enter')
time.sleep(5)
pyautogui.write('./configurechaincode.sh instantiate cli0 fabpki 1.0')
pyautogui.press('enter')
time.sleep(30)

#Inserir variáveis
pyautogui.write('cd fabpki-cli')
pyautogui.press('enter')
time.sleep(1)
pyautogui.write('python3 bancoLedger.py')
pyautogui.press('enter')
time.sleep(5)
pyautogui.write('python3 userLedger.py')
pyautogui.press('enter')
time.sleep(5)
for i in range(0, 9):
    pyautogui.write('python3 tra.py')
    pyautogui.press('enter')
    time.sleep(8)