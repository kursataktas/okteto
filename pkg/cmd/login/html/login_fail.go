// Copyright 2024 The Okteto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package html

import "html/template"

const (
	failTitle = `Okteto auth failed`

	failIcon = template.HTML(`<svg width="333" height="235" fill="none" xmlns="http://www.w3.org/2000/svg">
  <path fill-rule="evenodd" clip-rule="evenodd"
      d="M320.704 152.711C320.704 68.372 252.583 0 168.555 0 84.528 0 16.407 68.372 16.407 152.711c0 22.467 4.873 43.782 13.557 62.993h277.183c8.684-19.211 13.557-40.526 13.557-62.993Z"
      fill="#FFF0F0" />
  <path d="M17.37 215.704h297.556M1 215.704h9.63M322.63 215.704h9.63" stroke="#E54545" stroke-linecap="round"
      stroke-linejoin="round" />
  <g clip-rule="evenodd">
      <path fill-rule="evenodd"
          d="M246.4 201.259H90.711c-7.53 0-13.637-6.074-13.637-13.565V84.824c0-7.49 6.107-13.565 13.637-13.565H246.4c7.531 0 13.637 6.075 13.637 13.566v102.869c0 7.491-6.106 13.565-13.637 13.565Z"
          fill="#FFF0F0" />
      <path fill-rule="evenodd"
          d="M251.371 71.26H101.777c-13.64 0-24.703 6.967-24.703 15.558v94.219s73.452-22.418 89.462-54.889c16.011-32.47 84.835-54.889 84.835-54.889Z"
          fill="#fff" />
      <path fill-rule="evenodd"
          d="M194.26 113.653c-10.569-13.673-29.966-16.66-43.779-6.513-13.813 10.148-17.045 29.58-6.476 43.253 6.092 7.885 15.489 12.395 25.515 12.347 10.03-.047 18.999-4.773 24.769-12.724 1.863-2.676 1.601-6.998-1.414-8.705-3.014-1.707-6.182-1.148-8.329 1.81-4.813 6.633-13.281 9.326-21.276 6.817-8.004-2.511-13.169-9.752-13.342-17.945-.174-8.185 5.039-15.465 12.92-18.052 7.886-2.588 16.544.105 21.632 6.688 1.993 2.577 5.662 3.103 8.518 1.344 2.855-1.759 3.297-5.687 1.262-8.32Zm2.534 12.544a5.793 5.793 0 0 0-5.765-.446 5.68 5.68 0 0 0-3.203 4.778 5.863 5.863 0 0 0 2.565 5.2 5.795 5.795 0 0 0 5.766.446 5.68 5.68 0 0 0 3.202-4.778 5.86 5.86 0 0 0-2.565-5.2Z"
          fill="#E54545" />
      <path opacity=".579" fill-rule="evenodd"
          d="M251.371 71.26H101.777c-13.64 0-24.703 6.967-24.703 15.558v94.219s73.452-22.418 89.462-54.889c16.011-32.47 84.835-54.889 84.835-54.889Z"
          fill="#fff" />
      <path
          d="M246.4 201.259H90.711c-7.53 0-13.637-6.074-13.637-13.565V84.824c0-7.49 6.107-13.565 13.637-13.565H246.4c7.531 0 13.637 6.075 13.637 13.566v102.869c0 7.491-6.106 13.565-13.637 13.565Z"
          stroke="#E54545" stroke-linecap="round" stroke-linejoin="round" />
      <path fill-rule="evenodd" d="M183.963 224.37h-30.815l2.568-22.148h25.679l2.568 22.148Z" fill="#fff" />
      <path d="M183.963 224.37h-30.815l2.568-22.148h25.679l2.568 22.148Z" stroke="#E54545" stroke-linecap="round"
          stroke-linejoin="round" />
      <path fill-rule="evenodd"
          d="M202.259 234h-66.444c0-5.849 4.574-10.593 10.222-10.593h46c5.645 0 10.222 4.744 10.222 10.593ZM244.355 205.111c8.663 0 15.682-7.439 15.682-16.611v-5.537H77.074v5.537c0 9.172 7.02 16.611 15.683 16.611h151.598Z"
          fill="#fff" />
      <path
          d="M244.355 205.111c8.663 0 15.682-7.439 15.682-16.611v-5.537H77.074v5.537c0 9.172 7.02 16.611 15.683 16.611h151.598ZM202.259 234h-66.444c0-5.849 4.574-10.593 10.222-10.593h46c5.645 0 10.222 4.744 10.222 10.593Z"
          stroke="#E54545" stroke-linecap="round" stroke-linejoin="round" />
      <path d="M172.408 193.074a2.408 2.408 0 1 1-4.815-.001 2.408 2.408 0 0 1 4.815.001Z" stroke="#E54545"
          stroke-linecap="round" stroke-linejoin="round" />
  </g>
  <path fill-rule="evenodd" clip-rule="evenodd"
      d="M280.259 82.336c0 18.877-15.525 34.183-34.669 34.183-19.143 0-34.664-15.306-34.664-34.183 0-18.882 15.521-34.188 34.664-34.188 19.144 0 34.669 15.306 34.669 34.188Z"
      fill="#fff" />
  <path clip-rule="evenodd"
      d="M280.259 82.336c0 18.877-15.525 34.183-34.669 34.183-19.143 0-34.664-15.306-34.664-34.183 0-18.882 15.521-34.188 34.664-34.188 19.144 0 34.669 15.306 34.669 34.188Z"
      stroke="#E54545" stroke-linecap="round" stroke-linejoin="round" />
  <path fill-rule="evenodd" clip-rule="evenodd"
      d="M227.335 70.932c0-2.064.771-4.131 2.317-5.707 1.629-1.661 3.318-2.287 5.03-2.203a7.818 7.818 0 0 1 6.199 2.351l2.732 2.795c.852.793 1.694 1.624 2.522 2.45l5.082-4.96c3.153-3.097 8.262-3.097 11.419 0a7.814 7.814 0 0 1 0 11.199l-5.317 5.213 5.023 5.147c3.097 3.153 3.097 8.262 0 11.419a7.813 7.813 0 0 1-11.199 0l-5.214-5.317-5.155 5.018a8.097 8.097 0 0 1-4.902 2.274 8.203 8.203 0 0 1-.94.054 8.127 8.127 0 0 1-5.707-2.317c-1.661-1.63-2.287-3.318-2.203-5.03a7.823 7.823 0 0 1 2.351-6.2l2.795-2.73c.794-.853 1.625-1.696 2.45-2.524l-4.955-5.09a8.101 8.101 0 0 1-2.274-4.902 8.212 8.212 0 0 1-.054-.94Z"
      fill="#FFD6D6" />
  <path
      d="m229.652 65.225.357.35-.357-.35Zm-2.317 5.707h.5-.5Zm7.347-7.91-.025.5.032.001.031-.002-.038-.499Zm6.199 2.351.358-.35h-.001l-.357.35Zm2.732 2.795-.358.35.008.008.009.008.341-.366Zm2.522 2.45-.353.353.349.349.353-.345-.349-.358Zm5.082-4.96.349.357h.001l-.35-.358Zm11.419 0 .35-.357-.35.356Zm0 11.199.35.357-.35-.357Zm-5.317 5.213-.35-.357-.356.35.348.356.358-.35Zm5.023 5.147-.357.35h.001l.356-.35Zm0 11.419.357.35-.357-.35Zm-11.199 0-.357.35.357-.35Zm-5.214-5.317.357-.35-.348-.356-.357.347.348.359Zm-5.155 5.018-.349-.359-.001.001.35.358Zm-4.902 2.274-.048-.497-.009.001.057.496Zm-6.647-2.263.35-.357-.35.357Zm-2.203-5.03.5.025.001-.032-.002-.03-.499.037Zm2.351-6.2-.349-.357.349.358Zm2.795-2.73.35.357.008-.008.008-.009-.366-.34Zm2.45-2.524.354.354.348-.35-.343-.352-.359.348Zm-4.955-5.09.359-.349h-.001l-.358.349Zm-2.274-4.902.497-.048-.001-.009-.496.057Zm1.906-6.998a8.631 8.631 0 0 0-2.46 6.058h1c0-1.94.725-3.88 2.174-5.357l-.714-.7Zm5.411-2.351c-1.872-.092-3.695.602-5.411 2.351l.714.7c1.542-1.572 3.097-2.128 4.648-2.052l.049-.999Zm6.532 2.5a8.318 8.318 0 0 0-6.594-2.5l.076.998a7.32 7.32 0 0 1 5.804 2.202l.714-.7Zm2.732 2.796-2.731-2.795-.716.699 2.732 2.794.715-.698Zm2.518 2.444a101.775 101.775 0 0 0-2.535-2.461l-.681.732c.846.787 1.683 1.613 2.51 2.437l.706-.708Zm4.38-4.963-5.083 4.96.699.715 5.082-4.96-.698-.715Zm12.118 0c-3.352-3.286-8.772-3.287-12.119 0l.7.714c2.958-2.905 7.756-2.905 10.718 0l.701-.713Zm0 11.914a8.313 8.313 0 0 0 0-11.913l-.701.713a7.313 7.313 0 0 1 0 10.486l.701.714Zm-5.317 5.213 5.317-5.213-.701-.714-5.316 5.213.7.714Zm5.031 4.44-5.023-5.146-.716.698 5.024 5.147.715-.698Zm-.001 12.119c3.287-3.352 3.288-8.772 0-12.12l-.713.701c2.905 2.958 2.905 7.756 0 10.718l.713.7Zm-11.913 0a8.313 8.313 0 0 0 11.913 0l-.713-.7a7.314 7.314 0 0 1-10.486 0l-.714.7Zm-5.214-5.317 5.214 5.317.714-.7-5.214-5.317-.714.7Zm-4.449 5.026 5.155-5.018-.697-.717-5.156 5.018.698.717Zm-5.203 2.414a8.604 8.604 0 0 0 5.204-2.415l-.7-.715a7.602 7.602 0 0 1-4.6 2.135l.096.995Zm-.988.056c.333 0 .666-.019.997-.057l-.114-.993c-.293.033-.588.05-.883.05v1Zm-6.058-2.46a8.632 8.632 0 0 0 6.058 2.46v-1c-1.94 0-3.88-.725-5.357-2.174l-.701.714Zm-2.351-5.412c-.092 1.873.602 3.696 2.351 5.412l.701-.714c-1.573-1.543-2.13-3.097-2.053-4.648l-.999-.05Zm2.501-6.531a8.314 8.314 0 0 0-2.5 6.594l.997-.076a7.318 7.318 0 0 1 2.202-5.804l-.699-.714Zm2.795-2.732-2.795 2.731.699.716 2.795-2.732-.699-.715Zm2.445-2.519c-.826.829-1.662 1.676-2.462 2.536l.732.681c.788-.846 1.613-1.684 2.438-2.51l-.708-.707Zm-4.959-4.388 4.955 5.09.717-.697-4.955-5.09-.717.697Zm-2.414-5.203a8.602 8.602 0 0 0 2.415 5.204l.715-.7a7.603 7.603 0 0 1-2.135-4.6l-.995.096Zm-.056-.988c0 .333.019.666.057.997l.993-.114a7.907 7.907 0 0 1-.05-.883h-1Z"
      fill="#E54545" />
</svg>`)

	failContent = template.HTML(`<h1 class="error">Authorization failed</h1>
<h2>There was an error trying to login in.<br>Please try again later or use a different account.</h2>`)
)

func failData() *templateData {
	return &templateData{
		"Content": failContent,
		"Title":   failTitle,
		"Icon":    failIcon,
	}
}
