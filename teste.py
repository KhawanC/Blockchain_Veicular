import pyautogui, socket, time

#sk = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
#sk.connect(("8.8.8.8", 80))
#pyautogui.write(sk.getsockname()[0])

#Estabelecer conexão com a rede e subir o chaincode para a rede
pyautogui.hotkey('ctrl', 'shift', 't')
pyautogui.press('enter')
pyautogui.write('cd ..')
pyautogui.press('enter')
time.sleep(1)
pyautogui.write('docker-compose -f peer-orderer.yaml -f peer-ptb.yaml up -d')
pyautogui.press('enter')
time.sleep(23)
pyautogui.write('./configchannel.sh ptb.de -c')
pyautogui.press('enter')
time.sleep(10)
pyautogui.write('./configchaincode.sh install cli0 fabpki 1.0')
pyautogui.press('enter')
time.sleep(3)
pyautogui.write('./configchaincode.sh instantiate cli0 fabpki 1.0')
pyautogui.press('enter')
time.sleep(35)

#Inserir variáveis
pyautogui.write('cd fabpki-cli')
pyautogui.press('enter')
time.sleep(1)
pyautogui.write('python3 bancoLedger.py')
pyautogui.press('enter')
time.sleep(3)
pyautogui.write('python3 userLedger.py')
pyautogui.press('enter')
time.sleep(3)
for i in range(0, 9):
    pyautogui.write('python3 trajetoLedger.py')
    pyautogui.press('enter')
    time.sleep(4)

pyautogui.write('python3 calcularCreditos.py')
pyautogui.press('enter')