import React, { useState, useEffect } from 'react';
import { ResumoVeiculosCard } from '../../Components/ResumoVeiculosCard';
import { Api } from '../../Services/Api';
import { TableVeiculos } from '../../Components/TableVeiculos';
import { CardsResumo, MainBox } from './style'

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
      }, 5000)
    }, [ticker])

    const atualizarDados = async () => {
        const res = await Api.get(`veiculo`);
        setTotalVeiculos(e => res.data.length)
        const res2 = await Api.get(`fabricante`);
        setTotalFabricantes(e => res2.data.length)
        const res3 = await Api.get(`ordem`);
        setTotalTransacoes(e => res3.data.length)
        let acumulador = 0 
        for (let i = 0; i < res2.data.length; i++) {
          if(res2.data[i].Co2_Tot != 0) {
            acumulador += res2.data[i].Co2_Tot
          }
          if(res2.data[i].SaldoCarbono) {
            acumulador += res2.data[i].SaldoCarbono
          }
        }
        setTotalCarbono(e => acumulador)
    }

    return (
        <>  
          <MainBox>
            <CardsResumo>
              <ResumoVeiculosCard texto={'Veiculos registrados'} quantidade={totalVeiculos} img={'veic'}/>
              <ResumoVeiculosCard texto={'Fabricantes registrados'} quantidade={totalFabricantes} img={'fab'}/>
              <ResumoVeiculosCard texto={'Total de carbono emitido'} quantidade={totalCarbono} img={'carb'}/>
              <ResumoVeiculosCard texto={'Transações realizadas'} quantidade={totalTransacoes} img={'trans'}/>
            </CardsResumo>
            <TableVeiculos/> 
          </MainBox>
                       
        </>
    );
}