import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:frontend/common/widgets/custom_accent_button.dart';
import 'package:frontend/common/widgets/cutom_text_field.dart';
import 'package:frontend/common/widgets/main_wrapper.dart';
import 'package:frontend/features/user/presentation/add_staff/add_staff_controller.dart';
import 'package:frontend/styles/dark_theme.dart';
import 'package:frontend/styles/light_theme.dart';
import 'package:frontend/styles/text_styles.dart';
import 'package:frontend/utils/validator.dart';
import 'package:get/get.dart';

class AddStaffScreen extends StatelessWidget {
  AddStaffScreen({super.key});

  final controller = Get.put(AddStaffController());

  @override
  Widget build(BuildContext context) {
    return MainWrapper(
      content: SizedBox(
        height: MediaQuery.of(context).size.height,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisSize: MainAxisSize.max,
          children: [
            Wrap(
              children: [
                Image.asset(
                  'assets/logos/icon-vetlink.png',
                  width: 60.0,
                ),
                const SizedBox(
                  width: 10.0,
                ),
                SvgPicture.asset(
                  'assets/logos/logo-first-vetlink.svg',
                  semanticsLabel: 'VetLink Logo',
                  height: 60.0,
                ),
                SvgPicture.asset(
                  'assets/logos/logo-second-vetlink.svg',
                  semanticsLabel: 'VetLink Logo',
                  colorFilter: ColorFilter.mode(
                      Theme.of(context).brightness == Brightness.dark
                          ? darkAccentColor
                          : Colors.black,
                      BlendMode.srcIn),
                  height: 60.0,
                ),
              ],
            ),
            const SizedBox(
              height: 10.0,
            ),
            Container(
              padding:
                  const EdgeInsets.symmetric(vertical: 5.0, horizontal: 15.0),
              decoration: const BoxDecoration(
                color: lightSecondaryColor,
                borderRadius: BorderRadius.all(
                  Radius.circular(15.0),
                ),
              ),
              child: Text(
                'Add Staff',
                style: captionBoldPoppins.copyWith(color: Colors.white),
              ),
            ),
            const SizedBox(
              height: 15.0,
            ),
            Center(
              child: Card(
                child: ConstrainedBox(
                  constraints: const BoxConstraints(
                    maxWidth: 400,
                  ),
                  child: Padding(
                    padding: const EdgeInsets.all(15.0),
                    child: Form(
                      key: controller.formKey,
                      child: Column(
                        mainAxisSize: MainAxisSize.min,
                        children: [
                          Row(
                            children: [
                              Flexible(
                                child: CustomTextField(
                                  labelText: 'First Name',
                                  prefixIcon: const Icon(Icons.person),
                                  floatLabel: true,
                                  validator: Validator().notEmpty,
                                  controller: controller.firstNameField,
                                ),
                              ),
                              const SizedBox(
                                width: 15.0,
                              ),
                              Flexible(
                                child: CustomTextField(
                                  labelText: 'Last Name',
                                  prefixIcon: const Icon(Icons.person),
                                  floatLabel: true,
                                  validator: Validator().notEmpty,
                                  controller: controller.lastNameField,
                                ),
                              ),
                            ],
                          ),
                          const SizedBox(
                            height: 15.0,
                          ),
                          CustomTextField(
                            labelText: 'Email',
                            prefixIcon: const Icon(Icons.person),
                            floatLabel: true,
                            validator: Validator().email,
                            controller: controller.lastNameField,
                          ),
                          const SizedBox(
                            height: 20.0,
                          ),
                          Row(
                            children: [
                              Tooltip(
                                message: 'Veterinarian',
                                child: Row(
                                  mainAxisSize: MainAxisSize.min,
                                  children: [
                                    Obx(
                                      () => Radio(
                                        value: "VETERINARIAN STAFF",
                                        groupValue: controller.userRole.value,
                                        onChanged: (value) {
                                          controller.userRole.value = value!;
                                        },
                                      ),
                                    ),
                                    Container(
                                      width: 50,
                                      height: 50,
                                      padding: const EdgeInsets.all(12.0),
                                      decoration: BoxDecoration(
                                        color: Colors.grey[300],
                                        borderRadius:
                                            BorderRadius.circular(8.0),
                                      ),
                                      child: Tooltip(
                                        message: 'Veterinarian',
                                        child: Image.asset(
                                          'assets/icons/vet.png',
                                          fit: BoxFit.contain,
                                        ),
                                      ),
                                    ),
                                  ],
                                ),
                              ),
                              const SizedBox(
                                width: 10.0,
                              ),
                              Tooltip(
                                message: 'Receptionist',
                                child: Row(
                                  mainAxisSize: MainAxisSize.min,
                                  children: [
                                    Obx(
                                      () => Radio(
                                        value: "RECEPTIONIST STAFF",
                                        groupValue: controller.userRole.value,
                                        onChanged: (value) {
                                          controller.userRole.value = value!;
                                        },
                                      ),
                                    ),
                                    Container(
                                      width: 50,
                                      height: 50,
                                      padding: const EdgeInsets.all(12.0),
                                      decoration: BoxDecoration(
                                        color: Colors.grey[300],
                                        borderRadius:
                                            BorderRadius.circular(8.0),
                                      ),
                                      child: Image.asset(
                                        'assets/icons/receptionist.png',
                                        fit: BoxFit.contain,
                                      ),
                                    ),
                                  ],
                                ),
                              )
                            ],
                          ),
                          const SizedBox(
                            height: 15.0,
                          ),
                          Align(
                            alignment: Alignment.centerLeft,
                            child: Text(
                              'Using default password of "12345". \nPlease advise employees to change their passwords.',
                              style: captionRegularPoppins.copyWith(
                                fontStyle: FontStyle.italic,
                                color: lightOnSurfaceColor,
                              ),
                            ),
                          ),
                          const SizedBox(
                            height: 22.0,
                          ),
                          Align(
                            alignment: Alignment.centerRight,
                            child: SizedBox(
                              width: 150.0,
                              child: CustomAccentButton(
                                buttonLabel: 'Add Staff',
                                onPressed: () {
                                  controller.addStaff();
                                },
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
