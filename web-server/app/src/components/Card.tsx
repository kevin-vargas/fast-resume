import 'chart.js/auto';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import { Button } from '@mui/material';

type CardProps = {
    title: string,
    description: string,
    onClick: () => void,
}

export default ({ title, description, onClick }:CardProps) => {
    return (
        <Card> 
            <CardHeader  style={{ textAlign: 'center' }} title={title}/>
            <CardContent>
                <Typography style={{ textAlign: 'center' }} variant="body2">
                    {description}
                </Typography>
                <Button style={{ textAlign: 'center' }} onClick={onClick}>
                    Resume
                </Button>
            </CardContent>
        </Card>
    )
}