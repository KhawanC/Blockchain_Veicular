import React, { useState, useEffect } from 'react';
import { ResumoVeiculosCard } from '../../Components/ResumoVeiculosCard';
import { Api } from '../../Services/Api';
import { TableVeiculos } from '../../Components/TableVeiculos';
import { CardsResumo } from './style'

export const Main = (props) => {

    const [ticker, setTicker] = useState(0)
    const [totalVeiculos, setTotalVeiculos] = useState(0)
    const [totalFabricantes, setTotalFabricantes] = useState(0)
    const [totalCarbono, setTotalCarbono] = useState(0)
    const [totalTransacoes, setTotalTransacoes] = useState(0)

    useEffect(() => {

      setTimeout(function() {
        atualizarDados()
        setTicker(e => e + 1)
      }, 100)
    }, [ticker])

    const atualizarDados = async () => {
      try {
        const res = await Api.get(`veiculo`);
        setTotalVeiculos(e => res.data.length)
        const res2 = await Api.get(`fabricante`);
        setTotalFabricantes(e => res2.data.length)
        const res3 = await Api.get(`ordem`);
        setTotalTransacoes(e => res3.data.length)
        let acumulador = 0 
        for (let i = 0; i < res2.data.length; i++) {
          acumulador += res2.data[i].Co2_Tot 
        }
        setTotalCarbono(e => acumulador)
      } catch (error) {
        console.log(error)
      }
    }

    return (
        <>  
          <div className='h-screen bg-blue-800 pt-14'>
            <CardsResumo>
              <ResumoVeiculosCard texto={'Veiculos registrados'} quantidade={totalVeiculos} img={'veic'}/>
              <ResumoVeiculosCard texto={'Fabricantes registrados'} quantidade={totalFabricantes} img={'fab'}/>
              <ResumoVeiculosCard texto={'Total de carbono emitido'} quantidade={totalCarbono} img={'carb'}/>
              <ResumoVeiculosCard texto={'Transações realizadas'} quantidade={totalTransacoes} img={'trans'}/>
            </CardsResumo>
            <TableVeiculos/> 
          </div>
                       
        </>
    );
}