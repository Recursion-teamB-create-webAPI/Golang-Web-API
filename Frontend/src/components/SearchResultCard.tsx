import { Button, Card, CardBody, CardFooter, Heading, Image, Text } from '@chakra-ui/react'
import { Divider, Stack } from '@mui/material'

type SearchResultCardProps = {
  id: number,
  searchWord: string,
  imageURL: string,
}

const SearchResultCard = ({ id, searchWord, imageURL }: SearchResultCardProps) => {
  return (
    <>
      <Card>
        <CardBody >
          <Image
            src={imageURL}
            alt={`${searchWord} + ${id}`}
            borderRadius={'lg'}
          />
          <Stack mt='6' spacing='3'>
            <Heading size='md'>Search Word</Heading>
            <Text color='blue.600' fontSize='2xl'>
              {searchWord}
            </Text>
          </Stack>
        </CardBody>
        <Divider />
        <CardFooter>
          <Button variant='solid' colorScheme='blue'>
            Go to description
          </Button>
        </CardFooter>
      </Card >
    </>
  )
}

export default SearchResultCard
