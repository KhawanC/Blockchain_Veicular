import React from 'react';
import { Routes, Route } from 'react-router-dom'
import { Main } from "../Pages/Main"
import { Transacao } from "../Pages/Transacao"

export const Root = (params) => {
    return(
        <Routes>
            <Route path={'/'} element={<Main />}/>
            <Route path={'/transacao'} element={<Transacao/>}/>
        </Routes>
    )
};
