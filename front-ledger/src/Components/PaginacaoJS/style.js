import styled from 'styled-components'

export const Mainbox = styled.nav`

    display: flex;

    justify-content: center;
    align-items: center;

    ul {
        display: flex;

    }

    ul li {
        background-color: #3b51a1;

        display: flex;

        width: 30px;
        height: 30px;

        align-items: center;
        justify-content: center;

        font-size: 13px;
        border: 1px solid black;
    }

    ul li:hover {
        background-color: #89a0f5;
        cursor: pointer;
    }
`