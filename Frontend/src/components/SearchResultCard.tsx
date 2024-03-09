import {
  Card,
  CardBody,
  CardFooter,
  HStack,
  Image,
  VStack,
} from "@chakra-ui/react";
import { useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { v4 as uuidv4 } from "uuid";

type Props = {
  imageURL: string;
  totalResults: string[];
};

const SearchResultCard = ({ imageURL, totalResults }: Props) => {
  const [currentImage, setCurrentImage] = useState(imageURL);
  const navigate = useNavigate();
  const { username } = useParams();
  const handleDescription = () => {
    /*Go to each description page, but i don't know how rich information I can get by custom search json api.*/
    const descriptionURL = uuidv4();
    console.log(descriptionURL);
    navigate(`/${username}/description/${descriptionURL}`, {
      state: {
        imageURL: imageURL,
      },
    });
  };

  const handleGoPrevious = () => {
    const currentIndex = totalResults.indexOf(currentImage);
    const nextIndex =
      (currentIndex - 1 + totalResults.length) % totalResults.length;
    const nextImage = totalResults[nextIndex];
    setCurrentImage(nextImage);
  };

  const handleGoNext = () => {
    const currentIndex = totalResults.indexOf(currentImage);
    const nextIndex =
      (currentIndex + 1 + totalResults.length) % totalResults.length;
    const nextImage = totalResults[nextIndex];
    setCurrentImage(nextImage);
  };

  return (
    <>
      <VStack>
        <Card mb={"3"}>
          <CardBody>
            <VStack>
              <Image
                boxSize={"300px"}
                src={currentImage}
                alt={`Image: ${currentImage}`}
                borderRadius={"lg"}
              />
            </VStack>
          </CardBody>
          <CardFooter mt={8}>
            <HStack>
              <button
                className="bg-blue-500 rounded-xl p-3 text-white hover:bg-blue-600"
                onClick={handleGoPrevious}
              >
                {`\<\<  Previous`}
              </button>
              <button
                onClick={handleDescription}
                className="bg-blue-500 rounded-xl p-3 text-white hover:bg-blue-600"
              >
                Go to description
              </button>
              <button
                className="bg-blue-500 rounded-xl p-3 text-white hover:bg-blue-600"
                onClick={handleGoNext}
              >
                {`\>\>  Next`}
              </button>
            </HStack>
          </CardFooter>
        </Card>
      </VStack>
    </>
  );
};

export default SearchResultCard;
