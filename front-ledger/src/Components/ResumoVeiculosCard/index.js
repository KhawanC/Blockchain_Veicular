import React, { useState} from 'react';
import { Mainbox } from './style';

function ResumoVeiculosCard(props) {
    return (
        <>
            <Mainbox>
                <div className='linha1'>
                    <p>{props.texto}: {props.quantidade}</p>
                </div>
            </Mainbox>
        </>
    );
}

export default ResumoVeiculosCard;