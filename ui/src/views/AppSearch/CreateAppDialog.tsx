import React, { useEffect, useState } from "react";
import {
  Button,
  TextField,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Typography
} from "@mui/material";

export function CreateAppDialog(props: {
  open: boolean;
  setOpen: (b: boolean) => void;
}) {
  return (
    <Dialog maxWidth="md" fullWidth open={props.open} onClose={props.setOpen}>
      <DialogTitle>Create App</DialogTitle>
      <DialogContent>
        {/* <DialogContentText>
          To subscribe to this website, please enter your email address here. We
          will send updates occasionally.
        </DialogContentText> */}
        <TextField
          margin="dense"
          id="name"
          label="Name"
          fullWidth
          variant="standard"
        />
        <div style={{display: "flex", alignItems: "center"}}>
          <TextField
            margin="dense"
            id="repo"
            label="Repository"
            fullWidth
            variant="standard"
          />
          SELECT HERE
        </div>
      </DialogContent>
      <DialogActions>
        <Button onClick={() => props.setOpen(false)}>Cancel</Button>
        <Button onClick={() => props.setOpen(false)}>Subscribe</Button>
      </DialogActions>
    </Dialog>
  );
}
