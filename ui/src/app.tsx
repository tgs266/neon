import React from 'react';
import { Route, Routes } from 'react-router';
import { Home } from './views/Home';
import { ProductSearch } from './views/ProductSearch';
import './app.css'
import { Product } from './views/Product';

function App() {
    return (
        <div>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/products" element={<ProductSearch />} />
                <Route path="/products/:name" element={<Product />} />
            </Routes>
        </div>
    );
}
export default App;