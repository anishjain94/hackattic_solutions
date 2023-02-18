import cv2
import pytesseract
from PIL import Image

img = cv2.imread('modules/visual_basic_math/temp2.png', cv2.IMREAD_GRAYSCALE)

thresh = cv2.threshold(img, 100, 255, cv2.THRESH_BINARY_INV+cv2.THRESH_OTSU)[1]


# cv2.imwrite("modules/visual_basic_math/black.png", img)
# img = Image.open('modules/visual_basic_math/temp2.png')
print(pytesseract.image_to_string(img, config='--psm 6'))
