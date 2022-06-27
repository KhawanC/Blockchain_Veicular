import React from 'react';
import { useNavigate } from 'react-router-dom'
import { Navbar } from 'flowbite-react'
import { LinksBox, LogoBox, Mainbox } from './style';

export const Navvbar = (props) => {

    const navigate = useNavigate();

    return(
        <>
            <Mainbox className='bg-blue-800'>
                <LogoBox>
                    <img
                    src="https://i.imgur.com/DOGZ4Yd.png"
                    className="mr-5 h-10 sm:h-20"
                    alt="Carro"
                    />
                    <span className="self-center whitespace-nowrap text-4xl font-semibold">
                    Mobicrowd
                    </span>
                </LogoBox>
                <LinksBox>
                    <ul>
                        <li><span onClick={() => navigate('/')}>Inicio</span></li>
                        <li><span onClick={() => navigate('/transacao')}>Transações</span></li>
                    </ul>
                </LinksBox>
            </Mainbox>
        </>
    )
};
