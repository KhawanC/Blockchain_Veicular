import React from 'react';
import { Navbar } from 'flowbite-react'

export const Navvbar = (props) => {
    return(
        <>
            <Navbar fluid={true} rounded={true} className='bg-blue-500'>
            <Navbar.Brand href="./painel_administrativo">
                <img
                src="https://i.imgur.com/DOGZ4Yd.png"
                className="mr-5 h-10 sm:h-20"
                alt="Carro"
                />
                <span className="self-center whitespace-nowrap text-4xl font-semibold dark:text-white">
                Mobicrowd
                </span>
            </Navbar.Brand>
            <div className="flex md:order-2">
                <Navbar.Toggle />
            </div>
            <Navbar.Collapse>
                <Navbar.Link href="/"><span className='text-2xl text-black hover:text-white'>Inicio</span></Navbar.Link>
                <Navbar.Link href="/transacao"><span className='text-2xl text-black hover:text-white'>Transação</span></Navbar.Link>
            </Navbar.Collapse>
            </Navbar>

        </>
    )
};
