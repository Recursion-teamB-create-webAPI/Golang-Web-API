import {
  Card,
  CardBody,
  CardFooter,
  HStack,
  Image,
  VStack,
} from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { v4 as uuidv4 } from "uuid";
import { useSearchResultState } from "../store/SearchResultStore";

type Props = {
  imageURL: string;
  totalResults: string[];
};

const SearchResultCard = ({ imageURL, totalResults }: Props) => {
  const [currentImage, setCurrentImage] = useState(imageURL);
  const [results, setResults] = useState<string[]>([]);
  const navigate = useNavigate();
  const { username } = useParams();
  const [searchResults] = useSearchResultState((state) => [
    state.searchResults,
  ]);

  const handleDescription = () => {
    const descriptionURL = uuidv4();
    navigate(`/${username}/description/${descriptionURL}`, {
      state: {
        imageURL: imageURL,
      },
    });
  };

  const handleGoPrevious = () => {
    const currentIndex = results.indexOf(currentImage);
    const nextIndex = (currentIndex - 1 + results.length) % results.length;
    const nextImage = totalResults[nextIndex];
    setCurrentImage(nextImage);
  };

  const handleGoNext = () => {
    const currentIndex = results.indexOf(currentImage);
    const nextIndex = (currentIndex + 1 + results.length) % results.length;
    const nextImage = results[nextIndex];
    setCurrentImage(nextImage);
  };

  useEffect(() => {
    setResults(totalResults);
    setCurrentImage(totalResults[0]);
  }, [searchResults]);

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
