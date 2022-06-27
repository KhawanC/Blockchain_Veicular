import React from 'react';
import { MainBox } from './style';

export const PostTransJs = ({posts, loading}) => {
    if(loading) {
        return <h2>Loading...</h2>
    }

    return(
        <>  <MainBox>
                <table className='w-full text-left border-2 border-solid border-black'>
                    <thead>
                        <tr className='bg-blue-700'>
                            <th className='text-black'>Status Ordem</th>
                            <th className='text-black'>Proprietário</th>
                            <th className='text-black'>Saldo Ofertado</th>
                            <th className='text-black'>Tipo de Transação</th>
                            <th className='text-black'>Comprador</th>
                            <th className='text-black'>Valor do Ultimo Lance</th>
                        </tr>
                    </thead>
                    <tbody>
                        {posts.map(post => (
                            <tr key={post.id} className='border-2 border-solid border-black'>
                                <th scope='row' className='bg-blue-500'>
                                    {post.StatusOrdem}
                                </th>
                                <th scope='row'>
                                    {post.ProprietarioOrdem}
                                </th>
                                <th scope='row' className='bg-blue-500'>
                                    {post.SaldoOfertado}
                                </th>
                                <th scope='row'>
                                    {post.TipoTransacao}
                                </th>
                                <th scope='row'>
                                    {post.IdComprador}
                                </th>
                                <th scope='row'>
                                    {post.ValorUltimoLance}
                                </th>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </MainBox>
        </>
    )
};
