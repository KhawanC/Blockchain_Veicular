import React, { useState } from 'react';
import { TableTransacoes } from '../../Components/TableTransacoes';
import { FormTransacoes, ListTransacoes, MainBox } from './style';

export const Transacao = (params) => {
    const [iscriarOrdem, setCriarOrdem] = useState(true)
    const [isLancarOrdem, setLancarOrdem] = useState(false)
    const [isFinalizarLance, setFinalizarLance] = useState(false)

    return(
        <>
            <MainBox>
                <ListTransacoes>
                    <TableTransacoes>
                        
                    </TableTransacoes>
                </ListTransacoes>
                <FormTransacoes>
                    <>
                    </>
                </FormTransacoes>
            </MainBox>
        </>
    )
};
