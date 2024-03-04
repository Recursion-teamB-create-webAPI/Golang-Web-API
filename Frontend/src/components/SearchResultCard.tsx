import { Button, Card, CardBody, CardFooter, Image } from '@chakra-ui/react'
import { Divider, Stack } from '@mui/material'

const SearchResultCard = ({ imageURL }: SearchResult) => {
  const handleDescription = () => {
    /*Go to each description page, but i don't know how rich information I can get by custom search json api.*/
  }

  return (
    <>
      <Card mb={"3"}>
        <CardBody>
          <Image
            src={imageURL}
            alt={`Image: ${imageURL}`}
            borderRadius={'lg'}
          />
        </CardBody>
        <Divider />
        <CardFooter>
          <button
            onClick={handleDescription}
            className='bg-blue-500 rounded-xl p-3 text-white hover:bg-blue-600'
          >
            Go to description
          </button>
        </CardFooter>
      </Card >
    </>
  )
}

export default SearchResultCard
