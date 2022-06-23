import React, { useState, useEffect } from 'react';
import Navbar from '../../Components/Navbar';
import ResumoVeiculosCard from '../../Components/ResumoVeiculosCard';
import { Api } from '../../Services/Api';
import { CardsResumo } from './style';

function Main(props) {

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
            <Navbar/>
            <CardsResumo>
              <ResumoVeiculosCard texto={'Veiculos registrados'} quantidade={totalVeiculos}/>
              <ResumoVeiculosCard texto={'Fabricantes registrados'} quantidade={totalFabricantes}/>
              <ResumoVeiculosCard texto={'Total de carbono emitido'} quantidade={totalCarbono}/>
              <ResumoVeiculosCard texto={'Transações realizadas'} quantidade={totalTransacoes}/>
            </CardsResumo>
            
        </>
    );
}

export default Main;