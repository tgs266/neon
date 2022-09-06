import { Card, CardContent, Typography } from "@mui/material";
import { SxProps, Theme } from "@mui/system";
import React from "react";

export interface TitleCardProps {
    title: React.ReactNode 
    children?: React.ReactNode 
    sx?: SxProps<Theme>
}


export default function TitleCard(props: TitleCardProps) {
    return <Card sx={{padding: 2, ...props.sx}}>
            <Typography variant="h6" color="text.primary">
                {props.title}
            </Typography>
            {props.children}
    </Card>
}