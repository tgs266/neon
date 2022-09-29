import React, { useEffect, useState } from "react";
import {
  Button,
  TextField,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Typography,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
} from "@mui/material";
import { CredentialsService } from "../../services/CredentialsService";
import { Credentials } from "../../models/credentials";
import { CreateAppRequest } from "../../models/app";
import { AppService } from "../../services/AppService";
import { Product } from "../../models/product";
import { ProductService } from "../../services/ProductService";

export function AddProductDialog(props: {
  open: boolean;
  setOpen: (b: boolean) => void;
}) {

    const [products, setProducts] = useState<Product[]>([])
    const [selectedProduct, setSelectedProduct] = useState("")

    useEffect(() => {
        ProductService.listProducts(0, 0, "").then(r => {
            setProducts(r.data.items)
        })
    }, [])

    const save = () => {}

  return (
    <Dialog maxWidth="md" fullWidth open={props.open} onClose={props.setOpen}>
      <DialogTitle>Add Product</DialogTitle>
      <DialogContent>
          <FormControl margin="dense" variant="standard" fullWidth>
            <InputLabel id="products">Products</InputLabel>
            <Select
              labelId="products"
              value={selectedProduct}
              label="Products"
              onChange={(e) => setSelectedProduct(e.target.value)}
            >
              {products.map((c) => (
                <MenuItem key={c.name} value={c.name}>{c.name}</MenuItem>
              ))}
            </Select>
          </FormControl>
      </DialogContent>
      <DialogActions>
        <Button onClick={() => props.setOpen(false)}>Cancel</Button>
        <Button onClick={() => save()}>Add</Button>
      </DialogActions>
    </Dialog>
  );
}
