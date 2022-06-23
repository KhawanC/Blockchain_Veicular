import React from 'react';
import { Mainbox } from './style';
import Carro from "../../Source/icons/carro.png"

function Navbar(props) {
    return (
        <>
            <div>
                <Mainbox>
                    <div className='linha1'>
                        <div className='logo'>
                            <img src={Carro} alt="Carro"></img>
                            <p>MOBICROWD</p>
                        </div>
                        <div className='links'>
                            <ul>
                                <li>Resumo</li>
                                <li>Transações</li>
                                <li>Configurações</li>
                            </ul>
                        </div>
                    </div>
                </Mainbox>
            </div>   
        </>
    );
}

export default Navbar;