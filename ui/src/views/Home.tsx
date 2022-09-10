import React, { useEffect, useState } from 'react'
import Grid from '@mui/material/Grid';
import { Button, Card, CardActions, CardContent, Typography } from '@mui/material';
import { ProductService } from '../services/ProductService';
import { AppService } from '../services/AppService';
import { Link } from 'react-router-dom';

export function Home() {

    const [productCount, setProductCount] = useState<number>(null)
    const [appCount, setAppCount] = useState<number>(null)

    useEffect(() => {
        ProductService.listProducts(1, 0, "").then(r => {
            setProductCount(r.data.total)
        })
        AppService.listApps(1, 0, "").then(r => {
            setAppCount(r.data.total)
        })
    }, [])

    return <Grid container>
        <Grid item xs={6}>
            <Card>
                <CardContent>
                    <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
                        Apps
                    </Typography>
                    <Typography variant="h5" component="div">
                        {appCount != null ? `${appCount} Apps` : "SKELETON"}
                    </Typography>
                </CardContent>
                <CardActions>
                    <Button component={Link} to="/apps" size="small">View</Button>
                </CardActions>
            </Card>
        </Grid>
        <Grid item xs={6}>
            <Card>
                <CardContent>
                    <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
                        Products
                    </Typography>
                    <Typography variant="h5" component="div">
                        {productCount != null ? `${productCount} Products` : "SKELETON"}
                    </Typography>
                </CardContent>
                <CardActions>
                    <Button component={Link} to="/products" size="small">View</Button>
                </CardActions>
            </Card>
        </Grid>
    </Grid>
}