import React from 'react';
import { Mainbox } from './style';
import { CarOutline, LeafOutline, SettingsOutline, Wallet } from 'react-ionicons'


export const ResumoVeiculosCard = (props) => {
    return (
        <>
            <Mainbox className='bg-blue-700 border-2 border-solid rounded-3xl'>
                {props.img === 'trans' ? <>
                    <Wallet
                        height='25px'
                        width='25px'
                    />
                </> : ''}
                {props.img === 'veic' ? <>
                    <CarOutline
                        height='25px'
                        width='25px'
                    />
                </> : ''}
                {props.img === 'carb' ? <>
                    <LeafOutline
                        height='25px'
                        width='25px'
                    />
                </> : ''}
                {props.img === 'fab' ? <>
                    <SettingsOutline
                        height='25px'
                        width='25px'
                    />
                </> : ''}
                <div className='linha1'>
                    <p>{props.texto}</p>
                    <p>{props.quantidade}</p>
                </div>
            </Mainbox>
        </>
    );
}