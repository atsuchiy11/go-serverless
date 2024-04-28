FROM --platform=linux/arm64 golang:1.21 as build
WORKDIR /app
ARG HANDLER
# Copy dependencies list
COPY go.mod go.sum ./
# Copy path/to/handler/main.go from compose.yaml
#COPY $HANDLER .
COPY ./ .
# Build with optional lambda.norpc tag
RUN go build -tags lambda.norpc -o main $HANDLER

# Copy artifacts to a clean image
#FROM public.ecr.aws/lambda/provided:al2023 as runner
## Install aws-lambda-rie
#RUN curl -Lo /usr/bin/aws-lambda-rie https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie && chmod +x /usr/bin/aws-lambda-rie
#COPY --from=build /app/main ./main
#ENTRYPOINT [ "/usr/bin/aws-lambda-rie" ]
