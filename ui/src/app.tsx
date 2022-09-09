import React, { useState } from 'react';
import { Route, Routes } from 'react-router';
import { Home } from './views/Home';
import { ProductSearch } from './views/ProductSearch';
import './app.css'
import { Product } from './views/Product';
import { AppSearch } from './views/AppSearch';
import { Page } from './layout/Page';
import { createTheme } from '@mui/material/styles';
import { green, lightBlue } from '@mui/material/colors';
import { ThemeProvider } from "@mui/system";
import { App } from './views/App/App';
import { Install } from './views/Install/Install';

function MainApp() {

    const [mode, setMode] = useState<'light' | 'dark'>('light')

    const theme = createTheme({
        palette: {
            mode,
            primary: {
                main: green.A400,
            },
            secondary: {
                main: lightBlue[300],
            },
        },
    });


    return (
        <div style={{ height: "100vh" }}>
            <ThemeProvider theme={theme}>
                <Page mode={mode} setMode={setMode}>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/products" element={<ProductSearch />} />
                        <Route path="/products/:name" element={<Product />} />

                        <Route path="/apps" element={<AppSearch />} />
                        <Route path="/apps/:name" element={<App />} />
                        <Route path="/apps/:name/installs/:productName" element={<Install />} />
                    </Routes>
                </Page>
            </ThemeProvider>
        </div>
    );
}
export default MainApp;