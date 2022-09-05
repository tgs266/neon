import React, { useEffect, useState } from 'react'
import Grid from '@mui/material/Grid';
import { Button, Card, CardActions, CardContent, Table, TableBody, TableCell, TableHead, TableRow, TextField, Typography } from '@mui/material';
import { ProductService } from '../services/ProductService';
import { AppService } from '../services/AppService';
import { Product } from '../models/product';
import { Link } from 'react-router-dom';

export function ProductSearch() {

    const [name, setName] = useState<string>("")
    const [products, setProducts] = useState<Product[]>([])

    useEffect(() => {
        ProductService.listProducts(10, 0, name).then(r => {
            setProducts(r.data.items)
        })
    }, [name])

    return <div>
        <TextField
            id="outlined-basic"
            label="Name"
            variant="outlined"
            onChange={(e: React.ChangeEvent<HTMLInputElement>) => setName(e.target.value)}
            value={name}
        />

        <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead>
                <TableRow>
                    <TableCell>Name</TableCell>
                    <TableCell align="right"></TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {products.map((row) => (
                    <TableRow
                        key={row.name}
                        sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                    >
                        <TableCell component="th" scope="row">
                            {row.name}
                        </TableCell>
                        <TableCell align="right"><Button component={Link} to={`/products/${row.name}`}>View</Button></TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
    </div>
}