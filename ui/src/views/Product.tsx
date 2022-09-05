import React, { useEffect, useState } from 'react'
import Grid from '@mui/material/Grid';
import { Button, Card, CardActions, CardContent, Table, TableBody, TableCell, TableHead, TableRow, TextField, Typography } from '@mui/material';
import { ProductService } from '../services/ProductService';
import { AppService } from '../services/AppService';
import { Product } from '../models/product';
import { Link } from 'react-router-dom';
import { useParams } from 'react-router';

export function Product() {

    const [product, setProduct] = useState<Product>(null)
    const { name } = useParams();

    useEffect(() => {
        ProductService.get(name).then(r => {
            setProduct(r.data)
        })
    }, [name])

    return <div>
        {product?.name ? <div>{product.name}</div> : null}
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead>
                <TableRow>
                    <TableCell>Version</TableCell>
                    <TableCell align="right">Release Channel</TableCell>
                    <TableCell align="right">Helm Chart</TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {product?.releases && product.releases.map((row) => (
                    <TableRow
                        key={row.productVersion}
                        sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                    >
                        <TableCell component="th" scope="row">
                            {row.productVersion}
                        </TableCell>
                        <TableCell align="right" component="th" scope="row">
                            {row.releaseChannel}
                        </TableCell>
                        <TableCell align="right" component="th" scope="row">
                            {row.helmChart}
                        </TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
    </div>
}