/*
Copyright © 2022 António Inácio

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Stream type events are NOT meant to be subscribed by default all the time,
// i.e. during the entire application life cycle since boot time.
// The source brokers are only to be subscribed if there is an explicit
// consumer.
// Only one consumer is allowed, since we want to ensure delivery.
package stream
