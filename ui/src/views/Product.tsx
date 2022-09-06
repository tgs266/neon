import React, { useEffect, useState } from 'react'
import Grid from '@mui/material/Grid';
import { Button, Card, CardActions, CardContent, Chip, Table, TableBody, TableCell, TableHead, TableRow, TextField, Typography } from '@mui/material';
import { ProductService } from '../services/ProductService';
import { AppService } from '../services/AppService';
import { Product as ProductType } from '../models/product';
import { Link } from 'react-router-dom';
import { useParams } from 'react-router';
import TitleCard from '../components/DetailsTitleCard';
import Statistic from '../components/Statistic';
import { Box } from '@mui/system';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import ReleasesTable from '../components/Tables/ReleasesTable';
import Accordion from '../components/Accordion';
import InstallsTable from '../components/Tables/InstallsTable';


export function Product() {

    const [product, setProduct] = useState<ProductType>(null)
    const { name } = useParams();

    useEffect(() => {
        ProductService.get(name).then(r => {
            setProduct(r.data)
        })
    }, [name])

    return <div>
        <TitleCard title={<div style={{ display: "flex", alignItems: "center" }}>
            {product?.name ? <div>{product.name}</div> : null}
        </div>} sx={{ mb: 1 }}>
            <Box sx={{ mt: 1, display: "flex", alignItems: "center", gap: 2 }}>
                <Statistic label="Created At" value={new Date(product?.createdAt).toLocaleString()} />
                <Statistic label="Updated At" value={new Date(product?.updatedAt).toLocaleString()} />
            </Box>
        </TitleCard>
        <Card sx={{mb: 1}}>
            <Accordion title="Releases" defaultExpanded>
                <ReleasesTable product={product} />
            </Accordion>
        </Card>
        <Card>
            <Accordion title="Installs">
                <InstallsTable product={product} />
            </Accordion>
        </Card>
    </div>
}